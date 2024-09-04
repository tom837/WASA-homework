<template>
    <div v-if="isbanned">
        <button class="unban" @click="handleban">Unban</button>
    </div>
    <div v-else>
        <button class="ban" @click="handleban">Ban</button>
    </div>
</template>


<script>

    export default{
        data() {
            return {
                isbanned: this.Banning,
            };
        },
        props:{
            Banning:{
                type: Boolean,
                required : true
            }
        },
        methods: {
            handleban(){
                this.$emit('ban', { baned: this.isbanned});
                this.isbanned=!this.isbanned
            }
        },
    }
</script>

<style>

.ban{
    font-size: 25px;
    color: green;
}

.unban{
    font-size: 25px;
    color:red;
}



</style>
