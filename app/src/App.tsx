import React from 'react';
import { Admin, Resource, ListGuesser } from 'react-admin';
import jsonServerProvider from 'ra-data-json-server';

const dataProvider = jsonServerProvider("http://dev.admin.com:6868");

const App = () => (
  <Admin dataProvider={dataProvider}>
      <Resource name="users" list={ListGuesser} />
  </Admin>
)

export default App;
