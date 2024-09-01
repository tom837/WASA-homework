<template>
    <div>
        <form @submit.prevent="submitComment">
                    <div class="form-group">
                        <div class="comp">
                            <label for="comment">Comment</label>
                            <input
                                type="text"
                                v-model="new_comment"
                                class="form-control"
                                id="comment"
                                placeholder="Enter your comment"
                                required
                            />
                            <button type="submit" class="btn btn-primary btn-block">Comment</button>
                        </div>
                    </div>
        </form>
    </div>
</template>

<script>
    export default{
        data(){
            return {
                UserName: '',
                errorMessage: ''
            };
        },
        methods: {
            submitComment() {
                this.$emit('comment', {content: this.new_comment});
                this.errorMessage = '';
            }
        }
    };
</script>




<style>


</style>