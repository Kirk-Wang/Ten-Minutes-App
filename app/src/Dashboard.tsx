import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import Avatar from '@material-ui/core/Avatar';

const styles = {
  avatar: {
    width: 400,
    height: 400,
    borderRadius: 0
  },
};

export default withStyles(styles)(({ classes }) => (
    <Card>
        <CardHeader title="快速构建高质量的 React 应用" />
        <CardContent>
          <Avatar
            alt="GoLang"
            src="https://cdn-images-1.medium.com/max/1200/1*yh90bW8jL4f8pOTZTvbzqw.png" 
            className={classes.avatar}/>
        </CardContent>
    </Card>
))
