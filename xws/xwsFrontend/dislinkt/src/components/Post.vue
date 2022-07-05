<template>
  <div>
    <vs-card type="1" style="background: transparent">
      <template #title>
        <h3>{{ post.profile.firstName }} {{ post.profile.lastName }}</h3>
      </template>
      <template #img>
        <img :src="img" alt="">
      </template>
      <template #text>
        <p>
          {{ post.content.text }}
        </p>
      </template>
      <template #interactions v-if="checkIfLoggedUser()">
        <vs-button :active="checkIfFlat('CELEBRATE')" border icon @click="sendLike">
          <i class='bx bx-like'></i>
          <span v-if="likes > 0">
            {{ likes }}
          </span>
        </vs-button>
        <vs-button :active="checkIfFlat('DISLIKE')" danger border icon @click="sendDislike">
          <i class='bx bx-dislike'></i>
          <span v-if="dislikes > 0">
            {{ dislikes }}
          </span>
        </vs-button>
        <vs-button class="btn-chat" dark icon @click="openCommentDialog">
          <i class='bx bx-chat'></i>
          <span v-if="numOfComments > 0">
            {{ numOfComments }}
          </span>
        </vs-button>
        <vs-button v-if="ifUsersPost" danger icon @click="deletePost">
          <i class='bx bx-x'></i>
        </vs-button>
      </template>
    </vs-card>
    <vs-dialog :prevent-close="true" @close="resetCommentDialog" width="450px" v-model="commentDialog">
      <template #header>
        <h4 class="not-margin me-3 ms-3">
          Comments
        </h4>
      </template>
      <div class="con-form">
        <div v-for="c in comments" :key="c.id" class="mb-3">
          <div class="row" style="margin-bottom: 0px">
            <div class="col-6 d-flex justify-content-start" style="font-size: medium">
              <label :for="c.id"><b>{{ c.commentCreator.firstName }} {{ c.commentCreator.lastName }}</b></label>
            </div>
            <div class="col d-flex justify-content-end" style="font-size: small">
              <a disabled="">{{ formatDate(c.dateCreated) }}</a>
            </div>
          </div>
          <div class="row" style="padding-left: 15px; padding-right: 15px;">
            <div class="col-12 d-flex justify-content-center"  style="background-color: lavender; border-radius: 12px;">
                <p class="mt-3" style="font-size: medium">{{c.content}}</p>
            </div>
            <div v-if="currentUser(c)" class="col d-flex justify-content-end align-self-center">
              <div>
                <vs-button danger icon @click="deleteComment(c)">
                  <i class='bx bx-x'></i>
                </vs-button>
              </div>
            </div>
          </div>
        </div>
      </div>
      <template #footer>
        <div class="footer-dialog">
          <div class="con-form" style=" padding-top: 1rem">
            <div class="col mt-2" style="display: grid">
              <label for="textarea">Add new comment: </label>
              <textarea class="vs-input mb-2" style="width: 100%; border-radius: 12px" required primary v-model="newComment" id="textarea"></textarea>
              <vs-button block @click="addComment">
                Send
              </vs-button>
            </div>
          </div>
        </div>
      </template>
    </vs-dialog>
  </div>
</template>
<script>
import axios from "axios";
import moment from "moment";

export default {
  name: "Post",
  data() {
    return {
      img: "/proba.jpg",
      likes: 0,
      dislikes: 0,
      reactions: [],
      numOfComments: 0,
      userReaction: {},
      commentDialog: false,
      comments: [],
      newComment: '',
      user: null
    }
  },
  props: {
    post: {
      type: Object,
      required: true
    }
  },
  mounted() {
    this.getProfile()
    this.getReactions()
    this.getComments()
  },
  methods: {
    currentUser(comment){
      return comment.commentCreator.id === this.user.id
    },
    async getProfile() {
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/profile/${this.$store.getters.user?.id}`)
      if (response.data) {
        this.user = response.data.profile
      }
    },
    async getReactions() {
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/reaction/post/${this.post.id}`)
        .catch(e => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error loading reactions',
            color: 'danger',
            position: 'top-right'
          });
          throw e;
        })
      if (response.data.reactions.length > 0) {
        this.reactions = response.data.reactions
        this.likes = response.data.reactions.filter(r => r.type !== "DISLIKE").length
        this.dislikes = response.data.reactions.filter(r => r.type === "DISLIKE").length
      }
    },
    async getComments() {
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/comment/post/${this.post.id}`)
        .catch(e => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error loading comments',
            color: 'danger',
            position: 'top-right'
          });
          throw e;
        })
      if (response.data.comments?.length > 0) {
        this.numOfComments = response.data.comments.length
        this.comments = response.data.comments
        this.comments = this.comments.sort((a, b) => moment(a.dateCreated) - moment(b.dateCreated))
      }
    },
    checkIfLoggedUser() {
      return this.$store.getters.user?.role === "user"
    },
    async sendLike() {
      const userReaction = this.checkIfReacted()
      if (userReaction === null) {
        const response = await axios.post(`${process.env.VUE_APP_BACKEND}/reaction`, {
          id: null,
          userId: this.$store.getters.user.id,
          postId: this.post.id,
          type: "CELEBRATE",
          createdAt: moment().format()
        }).catch(e => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error sending reaction',
            color: 'danger',
            position: 'top-right'
          });
          throw e;
        })
        if (response.data) {
          await this.getReactions()
        }
      } else {
        if (userReaction.type === "CELEBRATE") {
          await axios.delete(`${process.env.VUE_APP_BACKEND}/reaction/${userReaction?.id}`).then(async () => {
            //this.likes -= 1
            await this.getReactions()
          })
          return
        }
        const response = await axios.post(`${process.env.VUE_APP_BACKEND}/reaction`, {
          id: userReaction.id,
          userId: userReaction.userId,
          postId: userReaction.postId,
          type: "CELEBRATE",
          createdAt: moment().format()
        }).catch(e => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error sending reaction',
            color: 'danger',
            position: 'top-right'
          });
          throw e;
        })
        if (response.data) {
          await this.getReactions()
        }
      }
    },
    async sendDislike() {
      const userReaction = this.checkIfReacted()
      if (userReaction === null) {
        const response = await axios.post(`${process.env.VUE_APP_BACKEND}/reaction`, {
          id: null,
          userId: this.$store.getters.user.id,
          postId: this.post.id,
          type: "DISLIKE",
          createdAt: moment().format()
        }).catch(e => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error sending reaction',
            color: 'danger',
            position: 'top-right'
          });
          throw e;
        })
        if (response.data) {
          await this.getReactions()
        }
      } else {
        if (userReaction.type === "DISLIKE") {
          await axios.delete(`${process.env.VUE_APP_BACKEND}/reaction/${userReaction.id}`).then(async () => {
            // if (userReaction.type === "CELEBRATE")
            //   await this.sendLike()
            await this.getReactions()
          })
          return
        }
        const response = await axios.post(`${process.env.VUE_APP_BACKEND}/reaction`, {
          id: userReaction.id,
          userId: userReaction.userId,
          postId: userReaction.postId,
          type: "DISLIKE",
          createdAt: moment().format()
        }).catch(e => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error sending reaction',
            color: 'danger',
            position: 'top-right'
          });
          throw e;
        })
        if (response.data) {
          await this.getReactions()
        }
      }
    },
    checkIfReacted() {
      for (let r of this.reactions) {
        if (r.userId === this.$store.getters.user.id && (r.type === "DISLIKE" || r.type === "CELEBRATE"))
          return r
      }
      return null
    },
    checkIfFlat(type) {
      for (let r of this.reactions) {
        if (r.userId === this.$store.getters.user.id && r.type === type)
          return true
      }
      return false
    },
    openCommentDialog() {
      this.commentDialog = true
    },
    async addComment() {
      if (this.newComment == '') return
      const comment = {
        content: this.newComment,
        commentCreator: {
          id: this.user.id,
          firstName: this.user.firstName,
          lastName: this.user.lastName,
        },
        postId: this.post.id,
        dateCreated: moment().format()
      }
      const response = await axios.post(`${process.env.VUE_APP_BACKEND}/comment`, comment)
        .catch(e => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error while posting a comment',
            color: 'danger',
            position: 'top-right'
          });
          throw e;
        })
      this.numOfComments += 1
      this.comments.push(response.data.comment)
      this.comments = this.comments.sort((a, b) => moment(a.dateCreated) - moment(b.dateCreated))
      this.resetCommentDialog()
    },
    resetCommentDialog() {
      this.newComment = ''
    },
    formatDate(date){
      return moment(date).format('lll')
    },
    async deleteComment(c){
      const response = await axios.delete(`${process.env.VUE_APP_BACKEND}/comment/${c.id}`)
        .catch(e => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error while deleting comment',
            color: 'danger',
            position: 'top-right'
          });
          throw e;
        })
      this.numOfComments -= 1
      this.comments = this.comments.filter(com => com.id !== c.id)
      this.comments = this.comments.sort((a, b) => moment(a.dateCreated) - moment(b.dateCreated))
      this.resetCommentDialog()
    },
    ifUsersPost(){
      return this.user.id === this.post.userId
    },
    async deletePost(){
      const response = await axios.delete(`${process.env.VUE_APP_BACKEND}/post/${this.post.id}`)
        .catch(e => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error while deleting post',
            color: 'danger',
            position: 'top-right'
          });
          throw e;
        })
      // this.$vs.notification({
      //   title: "Success",
      //   text: "Post deleted",
      //   color: "success",
      //   position: "top-right"
      // });
      this.$destroy();
      this.$el.parentNode.removeChild(this.$el);
    }
  }
}
</script>

<style scoped>

</style>
