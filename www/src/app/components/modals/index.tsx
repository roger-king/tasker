import React from 'react';
import { Box, Layer, Button, Heading } from 'grommet';
import { Close } from 'grommet-icons';

export interface BaseModalProps {
    setShowModal: any;
}

interface ModalProps extends BaseModalProps {
    header: string;
    width: any;
    onClickOutside: boolean;
    onEsc: boolean;
}

class Modal extends React.PureComponent<ModalProps, {}> {
    render(): JSX.Element {
        const { setShowModal, header, children, width, onClickOutside, onEsc } = this.props;

        return (
            <Layer
                modal
                onClickOutside={(): void => {
                    if (onClickOutside) {
                        setShowModal(false);
                    }
                }}
                onEsc={(): void => {
                    if (onEsc) {
                        setShowModal(false);
                    }
                }}
            >
                <Box width={width} pad="medium">
                    <Box direction="row">
                        <Button icon={<Close />} onClick={(): void => setShowModal(false)} />
                        <Heading level="3"> {header}</Heading>
                    </Box>
                    <Box margin="small">{children}</Box>
                </Box>
            </Layer>
        );
    }
}

export default Modal;
