<script>
    import comments from "@/components/comments.vue"
    import NewComment from "../components/NewComment.vue";
    import { getcomments } from '@/services/comments.js';
    import { getuserid } from '@/services/auth-store';
    import { getusername } from '@/services/auth-store';
    import { addcomment, removecomment } from "../services/comments";
    export default {
      components: { comments, NewComment },
      data() {
        return {
          comments: [],
          NewComment
        };
      },
      async mounted() {
        try {
            const photoid = this.$route.params.photoId;
            let comments= await getcomments(photoid); // feed() must return a promise
            this.username= await getusername();
            this.userid=await getuserid();
            this.comments = comments;
            console.log("comments",this.comments)
        } catch (error) {
          console.error('Error loading photos:', error);
        }
      },
      methods: {
		async PostComment(content) {
			this.errorMessage="";
			try {
				const response = await addcomment(this.userid,this.$route.params.photoId,content.content);
				console.log('comment added successfully', response);
                this.comments.push({
                    Id: response.Id,
                    Userid: this.userid,
                    User: this.username,
                    Comment: content.content,
                })
                console.log(comments)
			} catch (error) {
				this.errorMessage = error.message;
				console.error('Caught an error: '+ error.message);
			}
		},
        commented(comment){
          if (comment === this.userid){
            return true
          }else{
            return false
          }
        },

        async deletecomment(comment){
            try{
                let response= await removecomment(this.userid,comment.id);
                this.comments = this.comments.filter(c => c.Id !== comment.id);
                console.log('deleted comment successfuly', response);
                this.errorMessage='';
            } catch (error) {
                this.errorMessage = error.message;
                console.error('Caught an error: '+ error.message);
            }
	    }
    }
    }
</script>

<template>
    <div class="page-container">
        <div class="comments-container">
            <div v-if="comments===''">
                <p class="position">there are no comments</p>
            </div>
            <div v-else>
                <comments
                    v-for="comment in comments"
                    :key="comment.Id"
                    :commentid="comment.Id"
                    :username="comment.User"
                    :comment="comment.Comment"
                    :owner = "commented(comment.Userid)"
                    @deletecomment="deletecomment"
                />
            </div>
        </div>
        <div class="input-container">
            <NewComment @comment="PostComment" />
        </div>
    </div>
</template>



<style scoped>
.position {
    position: relative;
    top: 200px; /* Moves the paragraph 20px down from its normal position */
    text-align: center; /* Moves the paragraph 30px to the right from its normal position */
    font-size: 25px;

}

.input-container {
  padding: 10px;
  background-color: #f8f9fa;
  box-shadow: 0 -1px 5px rgba(0, 0, 0, 0.1);
}

.page-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.comments-container {
  flex: 1;
  overflow-y: auto;
  padding: 10px;
  box-sizing: border-box;
}


.comment-container {
  margin-bottom: 10px;
  padding: 10px;
  border-bottom: 1px solid #ccc;
}



</style>
