import React from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faBars } from '@fortawesome/free-solid-svg-icons'

class Breadcrumb extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      paths: [],
      sidebar: 1
    }
    this.hideSidebar = this.hideSidebar.bind(this);
  }

  componentDidMount() {
    let loc = sessionStorage.getItem('whereami');
    if (loc === null) {
      this.setState({paths: [{
        link: "http://localhost",
        name: "Home"
      }]})
    }
  }

  hideSidebar(e) {
    e.preventDefault();
    let status = this.state.sidebar;
    if (status === 1) {
      this.setState({ sidebar: 0});
    } else {
      this.setState({sidebar: 0});
    }
    console.log('sae');
  }

  render() {
    return <div id="breadcrumb" className="col-xs-12">
      <button onClick={this.hideSidebar} className="show-sidebar">
        <FontAwesomeIcon icon={faBars}/>
      </button>
      <ol className="breadcrumb pull-left">
        {this.state.paths.map(function(c, i) {
          return <li key={i}><a href={c.link}>{c.name}</a></li>
        })}
      </ol>
      <div id="social" className="pull-right">
        <span><i className="fa fa-google-plus"></i></span>
        <span><i className="fa fa-facebook"></i></span>
        <span><i className="fa fa-twitter"></i></span>
        <span><i className="fa fa-linkedin"></i></span>
        <span><i className="fa fa-youtube"></i></span>
      </div>
    </div>
  }
}

export default Breadcrumb;