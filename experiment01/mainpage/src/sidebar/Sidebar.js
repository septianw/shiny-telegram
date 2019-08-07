import React from 'react';

class Sidebar extends React.Component {
  constructor(props) {
    super(props);
    this.state = {menus: []};
  }

  componentDidMount() {
    let self = this;
    fetch('http://5d47faaf992ea9001444cd02.mockapi.io/api/v1/menus',{
      
    }).then(function (res) {
      return res.json();
    }).then(function (data) {
      self.setState({menus: data});
    })
  }

  componentWillUnmount() {
    // abis DidMount, jangan lupa Will Unmount
  }

  render() {
    return (
      <ul className="nav main-menu">
        {this.state.menus.map(function (c) {
          return <li key={c.id}>
            <a href={c.link} className="ajax-link">
              <i className="fa fa-{c.icon}"></i>
              <span className="hidden-xs">{c.name}</span>
            </a>
          </li>;
        })}
      </ul>
    );
  }
}

export default Sidebar;
