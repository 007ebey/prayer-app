import Login from "./pages/Login/Login";
import Dashboard from "./pages/Dashboard/Dashboard";

import {
  useState,
} from "react";

import { useAuth } from "./hooks";

import {
  useClerk,
} from "@clerk/clerk-react";

import type {
  LoginRole,
} from "./pages/Login/Login.types";
import { Container, Spinner } from "../primitives";

const App = () => {
  const {
    isLoaded,
    isSignedIn,
  } = useAuth();

  const {
    openSignIn,
  } = useClerk();

  
  const [loginRole, setLoginRole] =
    useState<LoginRole>("participant");

  const handleGoogleLogin = (
    role: LoginRole
  ) => {
    openSignIn({
      fallbackRedirectUrl: "/",
    });
  };

  const handleInstagramLogin = (
    role: LoginRole
  ) => {
    openSignIn({
      fallbackRedirectUrl: "/",
    });
  };


  /*
   * Clerk is still loading the authentication state.
   */
  if (!isLoaded) {
    return <Container
      className="flex items-center justify-center h-screen"
    >   
       <Spinner />
    </Container>;
  }

  /*
   * Clerk says the user is not authenticated.
   */
  if (!isSignedIn) {
    return (
      <Login
        role={loginRole}
        onRoleChange={setLoginRole}
        onGoogleLogin={handleGoogleLogin}
        onInstagramLogin={handleInstagramLogin}
      />
    );
  }

  /*
   * Temporary routing.
   *
   * Replace this later with the user's backend
   * profile and permissions.
   */
  return <Dashboard />;
};

export default App;
