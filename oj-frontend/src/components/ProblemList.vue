<!-- filepath: /oj-frontend/src/components/ProblemList.vue -->
<template>
  <div>
    <h1>Problem List</h1>
    <ul>
      <li v-for="problem in problems" :key="problem.id">{{ problem.title }}</li>
    </ul>
    <button @click="prevPage" :disabled="page === 1">Previous</button>
    <button @click="nextPage">Next</button>
  </div>
</template>

<script>
import { getProblems } from '../api/Problem';

export default {
  data() {
    return {
      problems: [],
      page: 1,
      pageSize: 10,
    };
  },
  methods: {
    async fetchProblems() {
      const data = await getProblems(this.page, this.pageSize);
      this.problems = data.items;
    },
    prevPage() {
      if (this.page > 1) {
        this.page--;
        this.fetchProblems();
      }
    },
    nextPage() {
      this.page++;
      this.fetchProblems();
    },
  },
  created() {
    this.fetchProblems();
  },
};
</script>

<style scoped>
/* 添加你的样式 */
</style>