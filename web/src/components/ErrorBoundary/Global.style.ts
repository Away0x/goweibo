import styled from 'styled-components';

const StyledGlobalErrorBoundary = styled.div`
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-color: #fff;
`;

const Inner = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  img {
    width: 340px;
  }
  h3 {
    color: #222;
    line-height: 30px;
    font-size: 19px;
  }
  p {
    color: #595959;
    font-size: 14px;
    line-height: 30px;
  }
`;

const Button = styled.div`
  width: 120px;
  height: 36px;
  line-height: 36px;
  text-align: center;
  border-radius: 17px;
  border: 1px solid ${({ theme }) => theme.themeColor};
  color: ${({ theme }) => theme.themeColor};
  margin-top: 20px;
  &:hover,
  &:active {
    color: rgba(${({ theme }) => theme.themeColor}, 0.8);
  }
`;

const Code = styled.pre`
  position: fixed;
  top: 0;
  padding-top: 20px;
  height: 50vh;
  overflow-y: auto;
  font-size: 10px;
`;

export default StyledGlobalErrorBoundary;

export { Inner, Button, Code };
