import apiClient from './apiClient';

export const getUsers = async (page, pageSize) => {
    const response = await apiClient.get('/users', {
        params: { page, pageSize },
    });
    return response.data;
};

export const getUserById = async (id) => {
    const response = await apiClient.get(`/users/${id}`);
    return response.data;
};

export const createUser = async (user) => {
    const response = await apiClient.post('/users', user);
    return response.data;
};

export const updateUser = async (id, user) => {
    const response = await apiClient.put(`/users/${id}`, user);
    return response.data;
};

export const deleteUser = async (id) => {
    const response = await apiClient.delete(`/users/${id}`);
    return response.data;
};