import React from 'react';
import { Admin, Resource } from 'react-admin';
import PostIcon from '@material-ui/icons/Book';
import { PostList, PostEdit, PostCreate } from './Posts';
import jsonServerProvider from 'ra-data-json-server';

const dataProvider = jsonServerProvider("http://dev.admin.com:6868");

const App = () => (
  <Admin dataProvider={dataProvider}>
      <Resource name="posts" list={PostList} edit={PostEdit} create={PostCreate} icon={PostIcon} />
  </Admin>
)

export default App;
