<script>
    import photo from '@/components/photo.vue';
    import { profile,deletepic } from '@/services/feed.js';
    import {like,unlike} from '@/services/like.js';
    import follow from '../components/follow.vue';
    import ban from '../components/ban.vue'
    import { Unfollow,Follow,getfollowers } from '../services/followers.js';
    import { getbanned,Ban,Unban } from '../services/banned.js';
	import { getuserid } from '../services/auth-store.js';
    export default {
      components: { photo, follow,ban},
      data() {
        return {
          photos: [],
          followerlist:[],
          bannedlist:[],
		  isbanned:[],
		  me:'',
          name: '',
          userid:''
        };
      },
      async mounted() {
        try {
          const userid = this.$route.params.id
          const name = this.$route.params.username
		  const me = await getuserid()
		  this.me=me
          this.name=name
          this.userid=userid
          let followerlist= await getfollowers();
          let bannedlist=await getbanned(this.me);
          this.bannedlist=bannedlist
		  let isbanned = await getbanned(this.userid)
		  this.isbanned=isbanned
          this.followerlist = followerlist
          let photos= await profile(userid);
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
        async handlefollow(content){
          if (content.following){
            try {
              const resp = await Unfollow(this.userid);
              console.log(resp);
              this.errorMessage='';
            } catch (error) {
              this.errorMessage = error.message;
              console.error('Caught an error: '+ error.message);
            }
          }else{
            try {

              const resp = await Follow(this.userid);
              console.log(resp);
              this.errorMessage='';
            } catch (error) {
              this.errorMessage = error.message;
              console.error('Caught an error: '+ error.message);
            }

          }
        },

        async handleban(content){
          if (content.baned){
            try {
              const resp = await Unban(this.userid);
              console.log(resp);
              this.errorMessage='';
            } catch (error) {
              this.errorMessage = error.message;
              console.error('Caught an error: '+ error.message);
            }
          }else{
            try {
              const resp = await Ban(this.userid);
              console.log(resp);
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
    <div class="page-container">
        <div class="photos-container">
			<div v-if="isbanned.includes(me)">
					<p class = "position">This user banned you</p>
					<div class="header">
						<h1>{{ name }} </h1>
						<div class="add">
							<div v-if="bannedlist.includes(userid)">
								<ban @ban="handleban" :Banning="true"/>
							</div>
							<div v-else>
								<ban @ban="handleban" :Banning="false"/>
							</div>
					</div>
				</div>
			</div>
			<div v-else>
            <div class="header">
                <h1>{{ name }} </h1>

                <div class="add">
                    <div v-if="followerlist.includes(userid)">
                        <follow @follow="handlefollow" :Following="true"/>
                    </div>
                    <div v-else>
                        <follow @follow="handlefollow" :Following="false"/>
                    </div>
                    <div v-if="bannedlist.includes(userid)">
                        <ban @ban="handleban" :Banning="true"/>
                    </div>
                    <div v-else>
                        <ban @ban="handleban" :Banning="false"/>
                    </div>

                </div>
            </div>
            <div v-if="photos==='No posts yet'">
                <p class="position">This user has not uploaded any pictures</p>
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
                    :owned = "false"
                    @like="HandleLikeEvent(photo.Id, $event.liked)"
                />
            </div>
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
