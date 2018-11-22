// -*- JavaScript -*-


class UserItem extends React.Component {
    render() {
      return (
        <tr>
          <td> {this.props.id}    </td>
          <td> {this.props.first} </td>
          <td> {this.props.last}  </td>
        </tr>
      );
    }
  }
  
  class UserList extends React.Component {
    constructor(props) {
      super(props);
      this.state = { user: [] };
    }
  
    componentDidMount() {
      this.serverRequest =
        axios
          .get("/users")
          .then((result) => {
             this.setState({ user: result.data });
          });
    }
  
    render() {
      const user = this.state.user.map((user, i) => {
        return (
          <UserItem key={i} id={user.Id} first={user.First} last={user.Last} />
        );
      });
  
      return (
        <div>
          <table><tbody>
            <tr><th>Id</th><th>First</th><th>Last</th></tr>
            {user}
          </tbody></table>
  
        </div>
      );
    }
  }
  
  ReactDOM.render( <UserList/>, document.querySelector("#root"));