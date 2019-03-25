import React from 'react';
import { Title } from 'react-admin';
import { withStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import Avatar from '@material-ui/core/Avatar';

const styles = () => ({
  content: {
    display: 'flex',
  },
  avatar: {
    margin: 10,
    width: 400,
    height: 400,
    borderRadius: 0,
  }
})

export default withStyles(styles, { withTheme: true })(({ classes }) => (
    <Card>
        <Title title="快速构建 Golang ❤️ MongoDB ❤️ React 应用" />
        <CardContent className={classes.content}>
          <Avatar
            alt="GoLang"
            src="https://cdn-images-1.medium.com/max/1200/1*yh90bW8jL4f8pOTZTvbzqw.png" 
            className={classes.avatar}/>
        </CardContent>
    </Card>
))
