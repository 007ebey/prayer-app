import {
    useState,
} from "react";

import Login from "./pages/Login/Login";
import Dashboard from "./pages/Dashboard/Dashboard";
import AdminPage from "./pages/Admin/AdminPage";

import type {
    LoginRole,
} from "./pages/Login/Login.types";

type AppView =
    | "login"
    | "participant"
    | "admin";

const App = () => {

    const [view, setView] =
        useState<AppView>("login");

    const [loginRole, setLoginRole] =
        useState<LoginRole>(
            "participant"
        );


    const handleLogin = (
        role: LoginRole
    ) => {

        // Temporary UI logic.
        //
        // Production:
        // authenticate provider
        //       ↓
        // backend verifies role
        //       ↓
        // navigate

        setView(role);

    };


    if (view === "admin") {

        return <AdminPage />;

    }


    if (view === "participant") {

        return <Dashboard />;

    }


    return (

        <Login
            role={loginRole}
            onRoleChange={
                setLoginRole
            }
            onGoogleLogin={
                handleLogin
            }
            onInstagramLogin={
                handleLogin
            }
        />

    );

};

export default App;