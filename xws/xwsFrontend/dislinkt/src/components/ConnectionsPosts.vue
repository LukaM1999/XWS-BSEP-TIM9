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
import axios from "axios";
import Post from "@/components/Post";

export default {
  name: "ConnectionsPosts",
  components: {Post},
  data() {
    return {
      posts: [],
      user: Object.assign({}, JSON.parse(localStorage.getItem('user'))),
    }
  },
  async mounted() {
    const loading = this.$vs.loading();
    const response = await axios.get(`https://localhost:8000/post/profile/62706d1b624b3da748f63fe3`).catch(error => {
      this.$vs.notification({
        title: 'Error',
        text: 'Error getting posts',
        color: 'danger',
        position: 'top-right'
      });
      loading.close();
      throw error;
    });
    if(response.data.posts.length > 0) {
      loading.close();
      this.posts = response.data.posts;
    }
  },

}
</script>

<style scoped>

</style>
