import React from 'react';
import { me } from '../data/auth';

export const UserContext = async (): Promise<React.Context<User>> => {
    const currentUser = await me();
    const userInfo = currentUser.data;
    return React.createContext<User>({ ...userInfo });
};
