<template>
    <div>
        <div class="user-container">
               <h1> <i class="bi bi-person icon-button" @click="userprofile">{{ username }}</i></h1>
        </div>
    </div>
</template>

<script>
    export default{
        props:{
            username :{
                type: String,
                required : true
            },
            id:{
                type: String,
                required:true
            }
        },
        data() {
            return {
            };
        },
        methods: {
            userprofile(){
                this.$router.push({name:"userprofile", params: { id: this.id, username: this.username }})
            }
        },
    };
</script>


<style>


.user-container{
    border: 1px solid #ccc;
    padding: 10px;
    justify-content: left;
    align-items: center;
    border-radius: 8px;
}


</style>







