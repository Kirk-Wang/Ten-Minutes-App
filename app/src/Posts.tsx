import React from "react";
import {
  List,
  Datagrid,
  TextField,
  ReferenceField,
  EditButton,
  Edit,
  SimpleForm,
  DisabledInput,
  ReferenceInput,
  SelectInput,
  TextInput,
  LongTextInput,
  Create,
  Filter,
  Responsive,
  SimpleList,
} from "react-admin";

const PostTitle = ({ record }:any) => {
  return <span>Post {record ? `"${record.title}"` : ''}</span>;
};

const PostFilter = (props: any) => (
    <Filter {...props}>
        <ReferenceInput label="User" source="userId" reference="users" alwaysOn>
            <SelectInput optionText="name" />
        </ReferenceInput>
    </Filter>
);

export const PostList = (props: any) => (
  <List {...props} filters={<PostFilter />}>
    <Datagrid>
      <TextField source="id" />
      <ReferenceField label="User" source="userId" reference="users">
          <TextField source="name" />
      </ReferenceField>
      <TextField source="title" />
      <TextField source="body" />
      <EditButton />
    </Datagrid>
  </List>
);

export const PostEdit = (props:any) => (
  <Edit title={<PostTitle />} {...props}>
    <SimpleForm>
      <DisabledInput source="id" />
      <ReferenceInput source="userId" reference="users">
        <SelectInput optionText="name" />
      </ReferenceInput>
      <TextInput source="title" />
      <LongTextInput source="body" />
    </SimpleForm>
  </Edit>
);

export const PostCreate = (props: any) => (
    <Create {...props}>
        <SimpleForm redirect="list">
            <ReferenceInput source="userId" reference="users">
                <SelectInput optionText="name" />
            </ReferenceInput>
            <TextInput source="title" />
            <LongTextInput source="body" />
        </SimpleForm>
    </Create>
);
