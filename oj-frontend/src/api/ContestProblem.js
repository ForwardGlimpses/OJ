import apiClient from './apiClient';

export const getContestProblems = async (contestProblemId, page, pageSize) => {
    const response = await apiClient.get(`/contest_problem/${contestProblemId}`, {
        params: { page, pageSize },
    });
    return response.data;
};

export const getContestProblemById = async (id) => {
    const response = await apiClient.get(`/contest_problem/${id}`);
    return response.data;
};

export const createContestProblem = async (contestProblem) => {
    const response = await apiClient.post('/contest_problem', contestProblem);
    return response.data;
};

export const updateContestProblem = async (id, contestProblem) => {
    const response = await apiClient.put(`/contest_problem/${id}`, contestProblem);
    return response.data;
};

export const deleteContestProblem = async (id) => {
    const response = await apiClient.delete(`/contest_problem/${id}`);
    return response.data;
};