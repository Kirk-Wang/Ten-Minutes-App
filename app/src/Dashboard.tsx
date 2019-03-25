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
  avatar1: {
    margin: 10,
    width: 200,
    height: 200,
    borderRadius: 0,
  },
  avatar2: {
    margin: 10,
    width: 150,
    height: 150,
    borderRadius: 0,
  },
  avatar3: {
    margin: 10,
    width: 180,
    height: 180,
    borderRadius: 0,
  },
})

export default withStyles(styles, { withTheme: true })(({ classes }) => (
    <Card>
        <Title title="快速构建高质量的 Golang ❤️ MongoDB ❤️ React 应用" />
        <CardContent className={classes.content}>
          <Avatar
            alt="GoLang"
            src="https://cdn-images-1.medium.com/max/1200/1*yh90bW8jL4f8pOTZTvbzqw.png" 
            className={classes.avatar1}/>
          <Avatar
            alt="React"
            src="https://raw.githubusercontent.com/github/explore/6c6508f34230f0ac0d49e847a326429eefbfc030/topics/react/react.png" 
            className={classes.avatar2}/>
          <Avatar
            alt="MongoDB"
            src="https://github.com/mongodb/mongo-go-driver/raw/master/etc/assets/mongo-gopher.png" 
            className={classes.avatar3}/>
        </CardContent>
    </Card>
))
