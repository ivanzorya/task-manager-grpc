import React, { useEffect } from "react";
import Alert from "react-bootstrap/Alert";

type TopAlertProps = {
    isVisible: boolean,
    setIsVisible: (val: boolean) => void,
    variant: string,
    message: string,
}

export const TopAlert = ({ isVisible, setIsVisible, variant, message }: TopAlertProps) => {

    const handleVisible = () => {
        setTimeout(() => {
            setIsVisible(false);
        }, 5000);
    }

    useEffect(() => {
        handleVisible();
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [isVisible])

    return (
        isVisible ?
            <Alert variant={variant}>
                {message}
            </Alert>
            :
            null
    )
}