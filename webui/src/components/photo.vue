<template>
    <div class="page-container">
        <div class="photo-container">
            <h1 class="username">{{ username }}</h1>
            <img :src="photoSrc" :alt="username + ' photo'" class="photo"/>
            <p> </p>
            <div class="actions">
                <div v-if="isLiked">
                    <i class="bi bi-heart-fill icon-button" @click="handleLike"></i>
                    <p>{{ numLikes }} Likes</p>
                </div>
                <div v-else>
                    <i class="bi bi-heart icon-button" @click="handleLike"></i>
                    <p>{{ numLikes }} Likes</p>
                </div>
                <div>
                    <i class="bi bi-chat icon-button" @click="showcomments"></i>
                    <p>{{ numcomments }} Comments</p>
                </div>
                <div v-if="owned">
                    <i class="bi bi-trash icon-button" @click="deletephoto"></i>
                    <p>Delete</p>
                </div>
            </div>
        </div>
    </div>
</template>



<script>
    export default{
        props:{
            username :{
                type: [String, null],
                required : true
            },
            photoBlob : {
                type : String,
                required : true
            },
            likes:{
                type : Number,
                required: true
            },
            liked:{
                type: Boolean,
                required: true
            },
            comments:{
                type: Number,
                required: true
            },
            photoId:{
                type: String,
                required: true
            },
            owned:{
                type:Boolean,
                required: false,
                default: false
            }
        },
        data() {
            return {
                photoSrc: '',
                isLiked: this.liked,
                numLikes: this.likes,
                numcomments: this.comments,
            };
        },
        methods: {
                updateLikesCount() {
                if (this.isLiked) {
                    this.numLikes += 1; // Increment likes
                } else {
                    this.numLikes -= 1; // Decrement likes
                }
            },
            handleLike() {
                this.$emit('like', { photo: this.photoId, liked: this.isLiked});
                this.isLiked=!this.isLiked
                this.updateLikesCount();
            },
            showcomments(){
                this.$router.push({ name: 'Comments', params: { photoId: this.photoId } });
            },
            deletephoto(){
                this.$emit('delphoto',{id: this.photoId})
            }
        },
        mounted() {
            // Create a URL for the binary data
            const decodedString = atob(this.photoBlob);
            const binaryDataStartIndex = decodedString.indexOf("image/jpeg") + "image/jpeg".length;
            const binaryDataStart = decodedString.indexOf("\r\n\r\n", binaryDataStartIndex) + 4;
            const jpegBinaryData = decodedString.substring(binaryDataStart);
            this.photoSrc = `data:image/jpeg;base64,${btoa(jpegBinaryData)}`;
        },
    };
</script>




<style scoped>

.page-container {
    display: flex;
    justify-content: center; /* Centers the content horizontally */
    align-items: center; /* Centers the content vertically */
    margin-top: 15px;
    margin-bottom: 5px;
}

.photo-container{
    display: inline-block;
    border: 1px solid #ccc;
    padding: 10px;
    justify-content: center;
    align-items: center;
    border-radius: 8px;
}

.photo{
    padding: 10px;
    text-align: center;
    max-width: 100%;
    height: 900PX;
    border-radius: 8px;
}

.username {
    margin-bottom: 10px;
    margin-left: 10px;
    text-align: left;
}


.icon-button {
  font-size: 25px;
  cursor: pointer;
  padding: 10px;
  border-radius: 50%; /* Makes it round */
}

.bi-heart-fill{
  color: red; /* Optional: Set the color of the heart icon */
  font-size: 25px;
  margin-right: 8px; /* Optional: Space between icon and text */
}

.bi-heart{
    color:red;
    font-size: 25px;
    margin-right: 8px;

}

.bi-chat{
    font-size: 25px;
    margin-left:8px;
}

.bi-trash {
  font-size: 25px;
  margin-left:8px;
}

.actions {
  display: flex;
  align-items: center;
  gap: 15px; /* Space between icons and text */
}

</style>