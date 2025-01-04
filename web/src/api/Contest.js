import apiClient from './apiClient';

export const getContests = async (page, pageSize) => {
    const response = await apiClient.get('/contest', {
        params: { page, pageSize },
    });
    return response.data;
};

export const getContestById = async (id) => {
    const response = await apiClient.get(`/contest/${id}`);
    return response.data;
};

export const createContest = async (contest) => {
    const response = await apiClient.post('/contest', contest);
    return response.data;
};

export const updateContest = async (id, contest) => {
    const response = await apiClient.put(`/contest/${id}`, contest);
    return response.data;
};

export const deleteContest = async (id) => {
    const response = await apiClient.delete(`/contest/${id}`);
    return response.data;
};