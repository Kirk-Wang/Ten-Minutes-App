import React from 'react';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import Avatar from '@material-ui/core/Avatar';

export default () => (
    <Card>
        <CardHeader title="快速构建高质量的 React 应用" />
        <CardContent>
          <Avatar alt="Material-UI" src="https://material-ui.com/static/images/material-ui-logo.svg" />
        </CardContent>
    </Card>
);
