<template>
  <vs-card type="1" style="background: transparent">
    <template #title>
      <h3>{{post.profile.firstName}} {{post.profile.lastName}}</h3>
    </template>
    <template #img>
      <img :src="img" alt="">
    </template>
    <template #text>
      <p>
        {{post.content.text}}
      </p>
    </template>
    <template #interactions>
      <vs-button icon>
        <i class='bx bx-like'></i>
        <span v-if="likes > 0">
          {{likes}}
        </span>
      </vs-button>
      <vs-button danger icon>
        <i class='bx bx-dislike'></i>
        <span v-if="dislikes > 0">
          {{dislikes}}
        </span>
      </vs-button>
      <vs-button class="btn-chat" dark icon>
        <i class='bx bx-chat'></i>
        <span v-if="comments > 0">
          {{comments}}
        </span>
      </vs-button>
    </template>
  </vs-card>
</template>

<script>
export default {
  name: "Post",
  data() {
    return {
      img: "/proba.jpg",
      likes: 0,
      dislikes: 0,
      comments: 0,
    }
  },
  props: {
    post: {
      type: Object,
      required: true
    }
  },
  mounted() {
  }
}
</script>

<style scoped>

</style>
