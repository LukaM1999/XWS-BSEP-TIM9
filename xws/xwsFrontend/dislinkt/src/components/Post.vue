<template>
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
      <vs-button class="btn-chat" dark icon>
        <i class='bx bx-chat'></i>
        <span v-if="comments > 0">
          {{ comments }}
        </span>
      </vs-button>
    </template>
  </vs-card>
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
      comments: 0,
      userReaction: {}
    }
  },
  props: {
    post: {
      type: Object,
      required: true
    }
  },
  mounted() {
    this.getReactions()
  },
  methods: {
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
        if(userReaction.type === "CELEBRATE"){
          await axios.delete(`${process.env.VUE_APP_BACKEND}/reaction/${userReaction?.id}`).then(async () => {
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
        if(userReaction.type === "DISLIKE"){
          await axios.delete(`${process.env.VUE_APP_BACKEND}/reaction/${userReaction.id}`).then(async () => {
            if(userReaction.type === "CELEBRATE")
              await this.sendLike()
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
    checkIfFlat(type){
      for (let r of this.reactions) {
        if (r.userId === this.$store.getters.user.id && r.type === type)
          return true
      }
      return false
    }
  }
}
</script>

<style scoped>

</style>
