<script>
    import photo from '@/components/photo.vue';
    import { profile,deletepic } from '@/services/feed.js';
    import {like,unlike} from '@/services/like.js'
    import { getuserid,getusername } from '@/services/auth-store';
    import add from "@/components/add.vue"
    export default {
      components: { photo, add},
      data() {
        return {
          photos: [],
		  name:''
        };
      },
      async mounted() {
        try {
          this.userid=await getuserid();
          let photos= await profile(this.userid); // feed() must return a promise
          this.name=await getusername();
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
        async Deletephoto(photo){
            try{
                let response = deletepic(photo.id)
                console.log("photo deleted successfully", response);
                this.photos = this.photos.filter(p => p.Id !== photo.id);
            }catch(error){
                console.error('Caught an error:'+error.message)
            }
        },

        async changename(){
            this.$router.push({ name: 'changename'})
        }
    },
  };



</script>


<template>
    <div class="page-container">
        <div class="photos-container">
            <div class="header">
                <h1>{{ name }}  <i class="bi bi-pen icon-button" @click="changename">change name</i>  </h1>
                <div class="add">
                    <add/>
                </div>
            </div>
            <div v-if="photos==='No posts yet'">
                <p class="position">Upload pictures so all your friends can see them</p>
            </div>
            <div v-else>
                <photo
                    v-for="photo in photos"
                    :key="photo.Id"
                    :photoId="photo.Id"
                    :username="null"
                    :photoBlob="photo.Photo"
                    :likes ="getLikesCount(photo.Likes)"
                    :liked= "Liked(photo.Likes)"
                    :comments = "getLikesCount(photo.Comments)"
                    :owned = "true"
                    @like="HandleLikeEvent(photo.Id, $event.liked)"
                    @delphoto="Deletephoto"
                />
            </div>
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

.add {
  padding: 10px;
  text-align: right;
  margin-right:25px;
}

.page-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
}
.header {
  display: flex;
  justify-content: space-between; /* Space between the name and add component */
  align-items: center; /* Center vertically */
  padding: 10px 25px; /* Add some padding */
  box-sizing: border-box;
}

.photos-container {
  flex: 1;
  overflow-y: auto;
  padding: 10px;
  box-sizing: border-box;
}


.photos-container {
  margin-bottom: 10px;
  padding: 10px;
  border-bottom: 1px solid #ccc;
}

.bi-pen{
    font-size: 15px;
}

</style>
