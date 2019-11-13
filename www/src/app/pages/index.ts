import React from 'react';

export const OverviewPage = React.lazy(() => import('./overview'));
export const TasksPage = React.lazy(() => import('./tasks'));
export const TaskPage = React.lazy(() => import('./task'));
export const SettingsPage = React.lazy(() => import('./settings'));
