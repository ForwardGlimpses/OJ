<template>
    <div>
        <h1>User List</h1>
        <ul>
            <li v-for="user in users" :key="user.id">{{ user.name }}</li>
        </ul>
        <button @click="prevPage" :disabled="page === 1">Previous</button>
        <button @click="nextPage">Next</button>
    </div>
</template>

<script>
import { getUsers } from '../api/User';

export default {
    data() {
        return {
            users: [],
            page: 1,
            pageSize: 10,
        };
    },
    methods: {
        async fetchUsers() {
            const data = await getUsers(this.page, this.pageSize);
            this.users = data.items;
        },
        prevPage() {
            if (this.page > 1) {
                this.page--;
                this.fetchUsers();
            }
        },
        nextPage() {
            this.page++;
            this.fetchUsers();
        },
    },
    created() {
        this.fetchUsers();
    },
};
</script>

<style scoped>
/* 添加你的样式 */
</style>