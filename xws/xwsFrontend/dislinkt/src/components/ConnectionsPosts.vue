<template>
  <div>
    <div class="row" style="margin-top: 7%">
      <div class="col">
        <div v-for="post in posts" :key="post.id">
          <div class="row">
            <div class="col-12 d-flex justify-content-center">
              <Post :post="post" class="mb-4" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Post from "@/components/Post";
import moment from "moment";

export default {
  name: "ConnectionsPosts",
  components: {Post},
  data() {
    return {
      posts: [],
      connections: [],
      user: null
    }
  },
  async mounted() {
    this.user = this.$store.getters.user
    await this.getConnections()
    if(this.connections.length > 0){
      for (const c of this.connections) {
        await this.getPosts(c.issuerId === this.user.id ? c.subjectId : c.issuerId)
      }
    }
  },
  methods: {
    async getPosts(id){
      const loading = this.$vs.loading();
      const response = await axios.get(`https://localhost:8000/post/profile/${id}`).catch(error => {
        this.$vs.notification({
          title: 'Error',
          text: 'Error getting posts',
          color: 'danger',
          position: 'top-right'
        });
        loading.close();
        throw error;
      });
      loading.close();
      if(!response.data?.posts) return;
      for (let post of response.data.posts) {
        this.posts.push(post)
      }
      this.sortPosts()
    },
    async getConnections(){
      const loading = this.$vs.loading();
      const response = await axios.get(`https://localhost:8000/connection/${this.user.id}`).catch(error => {
        this.$vs.notification({
          title: 'Error',
          text: 'Error getting connections',
          color: 'danger',
          position: 'top-right'
        });
        loading.close();
        throw error;
      });
      loading.close();
      this.connections = response.data.connections;
    },
    sortPosts(){
      this.posts = this.posts.sort((a, b) => moment(b.createdAt) - moment(a.createdAt))
    }
  }

}
</script>

<style scoped>

</style>
