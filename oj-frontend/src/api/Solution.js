import apiClient from './apiClient';

export const getSolutions = async (page, pageSize) => {
    const response = await apiClient.get('/solution', {
        params: { page, pageSize },
    });
    return response.data;
};

export const getSolutionById = async (id) => {
    const response = await apiClient.get(`/solution/${id}`);
    return response.data;
};

export const createSolution = async (solution) => {
    const response = await apiClient.post('/solution', solution);
    return response.data;
};

export const updateSolution = async (id, solution) => {
    const response = await apiClient.put(`/solution/${id}`, solution);
    return response.data;
};

export const deleteSolution = async (id) => {
    const response = await apiClient.delete(`/solution/${id}`);
    return response.data;
};