<script>
    import userlist from "@/components/userlist.vue"
    import { getuserid } from '@/services/auth-store';
    import { getuserlist } from "../services/feed";
    export default {
      components: { userlist },
      data() {
        return {
            users: [],
        };
      },
      async mounted() {
        try {
            let users= await getuserlist();
            let userid= await getuserid();
            this.userid=userid
            this.users = users;
            this.users = this.users.filter(u => u.ID !== userid);
        } catch (error) {
          console.error('Error loading photos:', error);
        }
      },
      methods: {
		},
    }
</script>

<template>
    <div class="page-container">
        <div class="comments-container">
            <div v-if="users===''">
                <p class="position">there are no other users</p>
            </div>
            <div v-else>
                <userlist
                    v-for="user in users"
                    :key="user.ID"
                    :id="user.ID"
                    :username="user.name"
                />
            </div>
        </div>
    </div>
</template>



<style scoped>
.position {
    position: relative;
    top: 200px; /* Moves the paragraph 20px down from its normal position */
    text-align: center; /* Moves the paragraph 30px to the right from its normal position */
    font-size: 25px;

}

.input-container {
  padding: 10px;
  background-color: #f8f9fa;
  box-shadow: 0 -1px 5px rgba(0, 0, 0, 0.1);
}

.page-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.comments-container {
  flex: 1;
  overflow-y: auto;
  padding: 10px;
  box-sizing: border-box;
}


.comment-container {
  margin-bottom: 10px;
  padding: 10px;
  border-bottom: 1px solid #ccc;
}



</style>
