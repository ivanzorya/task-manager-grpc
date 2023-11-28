import React, { useState } from "react";
import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import { clientUsers, clientAuth } from "../../client/client"
import { accessToken } from "../../App";
import { emailLogged } from "../../App";
import { useRecoilState } from "recoil";

type NavBarProps = {
    setAlertMessage: (val: string) => void,
    setAlertVariant: (val: string) => void,
    setIsAlertVisible: (val: boolean) => void,
}

export const NavBar = ({ setAlertMessage, setAlertVariant, setIsAlertVisible }: NavBarProps) => {
    const [token, setToken] = useRecoilState(accessToken);
    const [emailRecoiled, setEmailRecoiled] = useRecoilState(emailLogged);

    const [showLoginForm, setShowLoginForm] = useState<boolean>(false);

    const [showAccountCreationForm, setShowAccountCreationForm] = useState<boolean>(false);
    const [email, setEmail] = useState<string>("");
    const [password, setPassword] = useState<string>("");

    const addUser = () => {
        const user = {
            user: {
                id: "fake",
                email: email,
                password: password,
            }
        }
        clientUsers.queryCreateUser(user).then(
            res => {
                if (res.success) {
                    setShowLoginForm(true);
                    setShowAccountCreationForm(false);
                }
            }
        ).catch((e) => {
            setAlertMessage("User creation fail");
            setAlertVariant("danger");
            setIsAlertVisible(true);
        })
        setPassword("");
    }

    const logout = () => {
        setToken("");
        setEmail("");
        sessionStorage.setItem("accessToken", "");
        sessionStorage.setItem("email", "");
        setEmailRecoiled("")
    }

    const authUser = () => {
        const user = {
            email: email,
            password: password,
        }
        clientAuth.login(user).then((res) => {
            setToken(res.accessToken);
            setEmailRecoiled(email);
            sessionStorage.setItem("accessToken", res.accessToken);
            sessionStorage.setItem("email", email);
            setShowLoginForm(false);
        }).catch((e) => {
            setAlertMessage("Login fail");
            setAlertVariant("danger");
            setIsAlertVisible(true);

        })
        setPassword("");
    }

    return (
        <>
            <Navbar bg="primary" data-bs-theme="dark">
                <Container>
                    {
                        token !== "" ?
                            <>
                                <Navbar.Brand href="#home">Tasks</Navbar.Brand>
                                <Navbar.Collapse className="justify-content-end">
                                    <Col xs="auto">
                                        <Navbar.Text>
                                            Signed in as: <a href="#login">{emailRecoiled}</a>
                                        </Navbar.Text>
                                    </Col>
                                    <Col xs="auto">
                                        <Button
                                            onClick={() => {
                                                logout()
                                            }}
                                        >
                                            Logout
                                        </Button>
                                    </Col>
                                </Navbar.Collapse>

                            </>
                            :
                            showAccountCreationForm ?
                                <>
                                    <Navbar.Brand href="#home">New User</Navbar.Brand>
                                    <Form >
                                        <Row>
                                            <Col xs="auto">
                                                <Form.Control
                                                    value={email}
                                                    type="text"
                                                    placeholder="Email"
                                                    className=" mr-sm-2"
                                                    onChange={event => setEmail(event.target.value)}
                                                />
                                            </Col>
                                            <Col xs="auto">
                                                <Form.Control
                                                    value={password}
                                                    type="password"
                                                    placeholder="Password"
                                                    className=" mr-sm-2"
                                                    onChange={(event: React.ChangeEvent<HTMLInputElement>) => setPassword(event.target.value)}
                                                />
                                            </Col>
                                            <Col xs="auto">
                                                <Button
                                                    onClick={() => {
                                                        addUser()
                                                    }}
                                                >
                                                    Create Account
                                                </Button>
                                            </Col>
                                        </Row>
                                    </Form>
                                </>
                                :
                                showLoginForm ?
                                    <>
                                        <Navbar.Brand href="#home">Please login in</Navbar.Brand>
                                        <Form >
                                            <Row>
                                                <Col xs="auto">
                                                    <Form.Control
                                                        value={email}
                                                        type="text"
                                                        placeholder="Email"
                                                        className=" mr-sm-2"
                                                        onChange={event => setEmail(event.target.value)}
                                                    />
                                                </Col>
                                                <Col xs="auto">
                                                    <Form.Control
                                                        value={password}
                                                        type="password"
                                                        placeholder="Password"
                                                        className=" mr-sm-2"
                                                        onChange={(event: React.ChangeEvent<HTMLInputElement>) => setPassword(event.target.value)}
                                                    />
                                                </Col>
                                                <Col xs="auto">
                                                    <Button
                                                        onClick={() => {
                                                            authUser()
                                                        }}
                                                    >
                                                        Login
                                                    </Button>
                                                </Col>
                                            </Row>
                                        </Form>
                                    </>
                                    :
                                    <>
                                        <Navbar.Brand href="#home">Please login in</Navbar.Brand>
                                        <Form >
                                            <Row>

                                                <Col xs="auto">
                                                    <Button
                                                        type="submit"
                                                        onClick={() => setShowLoginForm(true)}
                                                    >
                                                        Sing in
                                                    </Button>
                                                </Col>
                                                <Col xs="auto">
                                                    <Button
                                                        type="submit"
                                                        onClick={() => setShowAccountCreationForm(true)}
                                                    >
                                                        Sing up
                                                    </Button>
                                                </Col>
                                            </Row>
                                        </Form>
                                    </>
                    }
                </Container>
            </Navbar>
        </>
    );
}

export default NavBar;