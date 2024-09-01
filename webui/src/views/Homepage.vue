<script>
    import handleLike from "@/components/photo.vue"
    import photo from '@/components/photo.vue';
    import { feed } from '@/services/feed.js';
    import {like,unlike} from '@/services/like.js'
    import { getuserid } from '@/services/auth-store';
    import showcomments from "@/components/photo.vue"
    export default {
      components: { photo, handleLike, showcomments},
      data() {
        return {
          photos: []
        };
      },
      async mounted() {
        try {
          let photos= await feed(); // feed() must return a promise
          this.userid=await getuserid();
          this.photos = photos;
        } catch (error) {
          console.error('Error loading photos:', error);
        }
      },
      methods: {
        getLikesCount(likes) {
          // If the Likes array contains only an empty string, return 0
          if (likes.length === 1 && likes[0] === '') {
            return 0;
          }
          // Otherwise, return the actual length of the Likes array
          return likes.length;
        },
        Liked(likes){
          if (likes.includes(this.userid)){
            return true
          }else{
            return false
          }
        },
        HandleLikeEvent(photoid, liked) {
            this.HandleLike(photoid, liked);
        },
        async HandleLike(photoid, liked){
          if (liked){
            try {
              const resp = await unlike(photoid);
              console.log('unliked photo successfuly', resp);
              this.errorMessage='';
            } catch (error) {
              this.errorMessage = error.message;
              console.error('Caught an error: '+ error.message);
            }
          }else{
            try {
              const resp = await like(photoid);
              console.log('liked photo successfuly', resp);
              this.errorMessage='';
            } catch (error) {
              this.errorMessage = error.message;
              console.error('Caught an error: '+ error.message);
            }

          }
        },
    },
  };



</script>

<template>
    <div>
      <div v-if="photos===''">
        <p class="position">Follow more users to see theire photos</p>
      </div>
      <div v-else>
        <photo
        v-for="photo in photos"
          :key="photo.Id"
          :username="photo.User"
          :photoBlob="photo.Photo"
          :likes ="getLikesCount(photo.Likes)"
          :liked= "Liked(photo.Likes)"
          :comments = "getLikesCount(photo.Comments)"
          @like="HandleLikeEvent(photo.Id, $event.liked)"
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
