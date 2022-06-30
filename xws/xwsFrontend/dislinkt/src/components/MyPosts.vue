<template>
  <div>
    <div class="row" style="margin-top: 7%">
      <div class="col d-flex justify-content-center">
        <div v-for="post in posts" :key="post.id">
          <Post :post="post" class="mb-4" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Post from "@/components/Post";
import axios from "axios";

export default {
  name: "MyPosts",
  components: {Post},
  data() {
    return {
      posts: [],
      user: null,
    }
  },
  async mounted() {
    this.user = this.$store.getters.user
    await this.getMyPosts()
    this.$parent.active = "myPosts"
  },
  methods: {
    async getMyPosts() {
      const loading = this.$vs.loading();
      const response = await axios.get(`https://localhost:8000/post/profile/${this.user.id}`).catch(error => {
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
      this.posts = response.data.posts;
    }
  }
}
</script>

<style scoped>

</style>
