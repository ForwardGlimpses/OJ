import { createStore } from 'vuex';

const store = createStore({
    state: {
        count: 0,
        user: null,
    },
    mutations: {
        increment(state) {
            state.count++;
        },
        setUser(state, user) {
            state.user = user;
        },
    },
    actions: {
        increment({ commit }) {
            commit('increment');
        },
        fetchUser({ commit }, userId) {
            // 模拟异步操作
            setTimeout(() => {
                const user = { id: userId, name: 'John Doe' };
                commit('setUser', user);
            }, 1000);
        },
    },
    getters: {
        doubleCount(state) {
            return state.count * 2;
        },
        userName(state) {
            return state.user ? state.user.name : '';
        },
    },
});

export default store;