export type Client = {
  name: string;
  clientId: string;
  createdAt: string;
  redirectUris: RedirectURI[];
};

export type RedirectURI = {
  id: string;
  uri: string;
  clientId: string;
  createdAt: string;
};

export type UpdateRedirectURI = {
  id: string;
  uri: string;
};