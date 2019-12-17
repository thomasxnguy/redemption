import React from "react";
import {
    CardContent,
    Card,
    CardHeader,
    TextField,
    Button
} from "@material-ui/core";

export default class extends React.Component {
    state = {
        bearerToken: "Default Value"
    };

    async componentDidMount() {
        const bearerToken = await localStorage.getItem("bearerToken");
        this.setState({ bearerToken });
    }

    handleChange = event => {
        this.setState({ bearerToken: event.target.value });
    };

    handleSubmit = event => {
        event.preventDefault();
        const { bearerToken } = this.state;
        localStorage.setItem("bearerToken", bearerToken);
    };

    render() {
        const { bearerToken } = this.state;
        return (
            <Card>
                <CardHeader title="Welcome to the Redemption Dashboard" />
                <CardContent>
                    <form onSubmit={this.handleSubmit}>
                        <TextField
                            id="bearer-token"
                            label="API Key"
                            onChange={this.handleChange}
                            value={bearerToken}
                            helperText="Bearer token for API authentication"
                            variant="outlined"
                        />
                        <Button
                            variant="contained"
                            color="primary"
                            type="submit"
                            size="large"
                        >
                            Save
                        </Button>
                    </form>
                </CardContent>
            </Card>
        );
    }
}
