import React from 'react';
import { Form, Input, Button } from 'antd';

import AuthContainer from 'containers/AuthContainer';

import StyledLogin from './style';

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};
const tailLayout = {
  wrapperCol: { offset: 8, span: 16 },
};

function Login() {
  const { loginAction } = AuthContainer.useContainer();

  const onFinish = (values: any) => {
    loginAction({
      username: values.password,
      password: values.password,
    });
  };

  return (
    <StyledLogin>
      <Form {...layout} name="basic" onFinish={onFinish}>
        <Form.Item
          label="Username"
          name="username"
          rules={[{ required: true, message: 'Please input your username!' }]}>
          <Input />
        </Form.Item>

        <Form.Item
          label="Password"
          name="password"
          rules={[{ required: true, message: 'Please input your password!' }]}>
          <Input.Password />
        </Form.Item>

        <Form.Item {...tailLayout}>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form>
    </StyledLogin>
  );
}

export default React.memo(Login);
