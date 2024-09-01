<template>
    <div>
        <div class="comment-container">
                <h1 class="username">{{ username }}</h1>
                <p style="font-size:30px">{{ comment }} </p>
            <div v-if="owner">
                <p></p>
                <i class="bi bi-trash icon-button" @click="deletecomment"></i>Delete
            </div>
        </div>
    </div>
</template>

<script>
    export default{
        props:{
            commentid:{
                type: String,
                required:true
            },
            username :{
                type: String,
                required : true
            },
            comment : {
                type : String,
                required : true
            },
            owner:{
                type: Boolean,
                required: true
            },
        },
        data() {
            return {
            };
        },
        methods: {
            deletecomment(){
                this.$emit('deletecomment', {id: this.commentid});
            }
        },
    };
</script>


<style>
.bi-trash {
  font-size: 25px;
  cursor: pointer;
  justify-content: right;
}

.comment-container{
    border: 1px solid #ccc;
    padding: 10px;
    justify-content: left;
    align-items: center;
    border-radius: 8px;
}


</style>







