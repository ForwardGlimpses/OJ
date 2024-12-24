import apiClient from './apiClient';

export const getSourceCodes = async (page, pageSize) => {
    const response = await apiClient.get('/source_code', {
        params: { page, pageSize },
    });
    return response.data;
};

export const getSourceCodeById = async (id) => {
    const response = await apiClient.get(`/source_code/${id}`);
    return response.data;
};

export const createSourceCode = async (sourceCode) => {
    const response = await apiClient.post('/source_code', sourceCode);
    return response.data;
};

export const updateSourceCode = async (id, sourceCode) => {
    const response = await apiClient.put(`/source_code/${id}`, sourceCode);
    return response.data;
};

export const deleteSourceCode = async (id) => {
    const response = await apiClient.delete(`/source_code/${id}`);
    return response.data;
};