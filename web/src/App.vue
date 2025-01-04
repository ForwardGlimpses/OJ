<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Your Vue.js App" />
  <div id="app">
    <ProblemList />
    <SolutionList />
    <SourceCodeList />
    <ContestList />
    <ContestProblemList />
    <UserList />
  </div>
</template>

<script>
import ContestList from './components/ContestList.vue';
import ContestProblemList from './components/ContestProblemList.vue';
import ProblemList from './components/ProblemList.vue';
import SolutionList from './components/SolutionList.vue';
import SourceCodeList from './components/SourceCodeList.vue';
import UserList from './components/UserList.vue';

export default {
  name: 'App',
  components: {
    ProblemList,
    SolutionList,
    SourceCodeList,
    ContestList,
    ContestProblemList,
    UserList,
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
