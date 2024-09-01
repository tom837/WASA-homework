<script>
    import comments from "@/components/comments.vue"
    import { getcomments } from '@/services/comments.js';
    import { getuserid } from '@/services/auth-store';
    export default {
      components: { comments },
      data() {
        return {
          comments: []
        };
      },
      async mounted() {
        try {
          let comments= await getcomments(); // feed() must return a promise
          this.userid=await getuserid();
          this.comments = comments;
          console.log("comments",this.comments)
        } catch (error) {
          console.error('Error loading photos:', error);
        }
      },
  };
</script>

<template>
    <div>
      <div v-if="comments===''">
        <p class="position">there are no comments</p>
      </div>
      <div v-else>
        <comments
        v-for="comment in comments"
          :key="comment.Id"
          :username="comment.User"
          :comment="comment.Comment"
        />
      </div>
    </div>
</template>



<style>
.position {
  position: relative;
    top: 200px; /* Moves the paragraph 20px down from its normal position */
    text-align: center; /* Moves the paragraph 30px to the right from its normal position */
    font-size: 25px;

}

</style>
