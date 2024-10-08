<script>
import userlist from "@/components/userlist.vue";
import { getuserid } from '@/services/auth-store';
import { getuserlist } from "../services/feed";

export default {
  components: { userlist },
  data() {
    return {
      users: [],
      searchQuery: '',
    };
  },
  async mounted() {
    try {
      let users = await getuserlist();
      let userid = await getuserid();
      this.userid = userid;
      this.users = users;
      this.users = this.users.filter(u => u.ID !== userid);
    } catch (error) {
      console.error('Error loading users:', error);
    }
  },
  computed: {
    filteredUsers() {
      return this.users.filter(user =>
        user.name.toLowerCase().includes(this.searchQuery.toLowerCase())
      );
    },
  },
};
</script>

<template>
  <div class="page-container">
    <div class="input-container">
      <!-- Search Bar -->
      <input 
        type="text" 
        v-model="searchQuery" 
        placeholder="Search users..." 
        class="search-input"
      />
    </div>

    <div class="comments-container">
      <div v-if="filteredUsers.length === 0">
        <p class="position">There are no other users</p>
      </div>
      <div v-else>
        <userlist
          v-for="user in filteredUsers"
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
  top: 200px;
  text-align: center;
  font-size: 25px;
}

.input-container {
  padding: 10px;
  background-color: #f8f9fa;
  box-shadow: 0 -1px 5px rgba(0, 0, 0, 0.1);
}

.search-input {
  width: 100%;
  padding: 10px;
  font-size: 16px;
  box-sizing: border-box;
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
