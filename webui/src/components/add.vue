<template>
    <div>
        <i class="bi bi-plus icon-button" @click="uploadphoto"></i>
        <p>upload photo</p>
    </div>
</template>


<script>

    export default{
        data() {
            return {
            };
        },
        methods: {
            uploadphoto(){
                this.$router.push({ name: 'upload'});
            }
        },
    }
</script>

<style>
.bi-plus{
    font-size:25px;
}

.icon-button {
  font-size: 25px;
  cursor: pointer;
  padding: 10px;
  border-radius: 50%; /* Makes it round */
}

</style>
