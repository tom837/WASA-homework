<template>
    <div v-if="isfollowing">
        <button class="unfollow" @click="handlefollow">Unfollow</button>
    </div>
    <div v-else>
        <button class="follow" @click="handlefollow">Follow</button>
    </div>
</template>


<script>

    export default{
        data() {
            return {
                isfollowing: this.Following,
            };
        },
        props:{
            Following:{
                type: Boolean,
                required : true
            }
        },
        methods: {
            handlefollow(){
                this.$emit('follow', { following: this.isfollowing});
                this.isfollowing=!this.isfollowing
            }
        },
    }
</script>

<style>

.follow{
    font-size: 25px;
    color: green;
}

.unfollow{
    font-size: 25px;
    color:red;
}



</style>
