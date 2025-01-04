<!-- filepath: /C:/Users/乔书祥/Desktop/OJ/oj-frontend/src/components/SourceCodeList.vue -->
<template>
    <div>
        <h1>Source Code List</h1>
        <ul>
            <li v-for="sourceCode in sourceCodes" :key="sourceCode.id">{{ sourceCode.title }}</li>
        </ul>
        <button @click="prevPage" :disabled="page === 1">Previous</button>
        <button @click="nextPage">Next</button>
    </div>
</template>

<script>
import { getSourceCodes } from '../api/SourceCode';

export default {
    data() {
        return {
            sourceCodes: [],
            page: 1,
            pageSize: 10,
        };
    },
    methods: {
        async fetchSourceCodes() {
            const data = await getSourceCodes(this.page, this.pageSize);
            this.sourceCodes = data.items;
        },
        prevPage() {
            if (this.page > 1) {
                this.page--;
                this.fetchSourceCodes();
            }
        },
        nextPage() {
            this.page++;
            this.fetchSourceCodes();
        },
    },
    created() {
        this.fetchSourceCodes();
    },
};
</script>

<style scoped>
/* 添加你的样式 */
</style>