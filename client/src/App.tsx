import React, { useState, useEffect } from "react";
import "./App.css";
import { Tasks } from "./features/tasks/tasks"
import "bootstrap/dist/css/bootstrap.min.css";
import { NavBar } from "./features/tasks/navbar";
import { atom, useRecoilState } from "recoil";
import { TopAlert } from "./features/tasks/topAlert"

export const accessToken = atom({
  key: "accessToken",
  default: '',
});

export const emailLogged = atom({
  key: 'emailLogged',
  default: '',
});

const App = () => {
  const [token, setToken] = useRecoilState(accessToken);
  const [emailRecoiled, setEmailRecoiled] = useRecoilState(emailLogged);

  const [isAlertVisible, setIsAlertVisible] = useState<boolean>(false);
  const [alertMessage, setAlertMessage] = useState<string>("");
  const [alertVariant, setAlertVariant] = useState<string>("");

  useEffect(() => {
    if (token === "") {
      const sessionToken = sessionStorage.getItem("accessToken")
      if (sessionToken) {
        setToken(sessionToken)
      }
    }

    if (emailRecoiled === "") {
      const sessionEmail = sessionStorage.getItem("email")
      if (sessionEmail) {
        setEmailRecoiled(sessionEmail)
      }
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <div>
      <TopAlert
        isVisible={isAlertVisible}
        setIsVisible={setIsAlertVisible}
        variant={alertVariant}
        message={alertMessage}
      />
      <NavBar
        setAlertMessage={setAlertMessage}
        setAlertVariant={setAlertVariant}
        setIsAlertVisible={setIsAlertVisible}
      />
      <Tasks
        setAlertMessage={setAlertMessage}
        setAlertVariant={setAlertVariant}
        setIsAlertVisible={setIsAlertVisible}
      />
    </div>
  )
}

export default App;
