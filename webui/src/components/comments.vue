<template>
    <div class="page-container">
        <div class="comment-container">
            <h1 class="username">{{ username }}</h1>
            <p>{{ comment }} </p>
        </div>
        <form @submit.prevent="submitComment">
                    <div class="form-group">
                        <label for="comment">Comment</label>
                            <input
                                type="text"
                                v-model="new_comment"
                                class="form-control"
                                id="comment"
                                placeholder="Enter your comment"
                                required
                            />
                    </div>
                <button type="submit" class="btn btn-primary btn-block">Comment</button>
        </form>
    </div>
</template>

<script>
    export default{
        props:{
            username :{
                type: String,
                required : true
            },
            Comment : {
                type : String,
                required : true
            },
        },
        data() {
            return {
                Comment: '',
            };
        },
        methods: {
            submitComment() {
                this.$emit('comment', { photo: this.new_comment});
            }
        },
    };
</script>








