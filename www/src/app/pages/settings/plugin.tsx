import React, { useState, useEffect } from 'react';
import { Box, Heading } from 'grommet';
import { listSettings } from '../../data/settings';

const PluginSettingPage: React.FC = () => {
    const [loading, setLoading] = useState<boolean>(true);
    const [settings, setSettings] = useState<Setting[] | null>(null);
    const [error, setError] = useState<any>();

    useEffect(() => {
        listSettings()
            .then(({ data }: { data: Setting[] | null }) => {
                setLoading(false);
                setSettings(data);
            })
            .catch((err: any) => {
                setError(err);
            });
    }, []);

    if (error) {
        return <Box>{error}</Box>;
    }

    if (loading) {
        return <Box>loading...</Box>;
    }

    if (settings && settings.length > 0) {
        return (
            <Box>
                <Heading level="4">Repositories:</Heading>
                <Box>
                    {settings.map((s: Setting) => {
                        return <pre key={`${s.repo_name}-code`}>{s.repo_name}</pre>;
                    })}
                </Box>
            </Box>
        );
    }

    return <Box>configure your plugin</Box>;
};

export default PluginSettingPage;
