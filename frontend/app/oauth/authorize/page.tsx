"use client";

import { useEffect, useState } from "react";
import useAxios from "@/hooks/useAxios";
import LoginView from "@/components/oauth/authorize/LoginView";
import ConsentView from "@/components/oauth/authorize/ConsentView";

export default function AuthorizePage() {
  const axios = useAxios();

  const [loading, setLoading] = useState(true);
  const [loggedIn, setLoggedIn] = useState(true);
  const [client, setClient] = useState(null);

  useEffect(() => {
    checkStatus();
  }, []);

  async function checkStatus() {
    try {
      const res = await axios.get("/oauth/authorize");
      setLoggedIn(res.data.logged_in);
      setClient(res.data.client);
    } catch (e) {
      console.error(e);
    } finally {
      setLoading(false);
    }
  }

  if (loading) return <div>Loading...</div>;

  if (!loggedIn) {
    return <LoginView />;
  }

  return <ConsentView />;
}
