import apiClient from './apiClient';

export const getProblems = async (page, pageSize) => {
    const response = await apiClient.get('/problem', {
        params: { page, pageSize },
    });
    return response.data;
};

export const getProblemById = async (id) => {
    const response = await apiClient.get(`/problem/${id}`);
    return response.data;
};

export const createProblem = async (problem) => {
    const response = await apiClient.post('/problem', problem);
    return response.data;
};

export const updateProblem = async (id, problem) => {
    const response = await apiClient.put(`/problem/${id}`, problem);
    return response.data;
};

export const deleteProblem = async (id) => {
    const response = await apiClient.delete(`/problem/${id}`);
    return response.data;
};