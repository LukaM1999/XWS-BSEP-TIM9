<template>
  <div style="padding: 6rem">
    <div class="row">
          <div class="col-4 mt-4" v-for="comment in comments" :key="comment.id">
            <vs-card style="padding-left: 5%;">
              <template #title>
                <div class="row">
                  <div class="col">
                    <p class="text-lg-end">{{comment.dateCreated | formatDate}} </p>
                    <h3 class="text-start" style="font-weight: bold; font-size: x-large">{{comment.subject}}</h3>
                  </div>
                </div>
                <div class="row mt-1">
                  <div class="col">
                    <star-rating :increment="0.5"
                                 :max-rating="5"
                                 :star-size="25"
                                 :read-only="true"
                                 :rating="comment.rating"
                                 :round-start-rating="false">
                    </star-rating>
                  </div>
                </div>
              </template>
              <template #text>
                <div class="row mt-3 mb-4">
                  <div class="col">
                    <p class="text-lg-start" style="font-size: large">{{comment.content}}</p>
                  </div>
                </div>
                <div class="row">
                  <div class="col">
                    <p class="text-lg-end"><i>{{comment.position}}</i></p>
                  </div>
                </div>
                <div class="row">
                  <div class="col">
                    <p class="text-lg-end"><i>{{comment.engagement}}
                      ({{comment.currentlyEmployed ? 'Currently employed': 'Left company'}})</i></p>
                  </div>
                </div>
              </template>
            </vs-card>
          </div>
        </div>
  </div>
</template>

<script>
import axios from "axios";
import StarRating from "vue-star-rating";

export default {
  name: "Comments",
  data() {
    return {
      comments: [],
    }
  },
  components:{
    StarRating,
  },
  mounted() {
    this.$parent.active = 'comments';
    this.getComments();
  },
  methods: {
    async getComments() {
      await axios.get(`${process.env.VUE_APP_BACKEND}/company/${this.$route.params.companyName}/comment`).then(response => {
        this.comments = response.data;
      });
    },
  },
}
</script>

<style scoped>

</style>
