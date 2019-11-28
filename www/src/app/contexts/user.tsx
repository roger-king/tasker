import React from 'react';

export const UserContext = React.createContext<User>({ name: '', email: '', githubURL: '', username: '', bio: '' });
