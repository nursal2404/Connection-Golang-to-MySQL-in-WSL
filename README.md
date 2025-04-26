# Setting up MySQL User for Golang Application on WSL

This guide explains how to create a MySQL user (`mysql_wsl_username` and `mysql_wsl_password`) for use in a Golang application running on WSL (Windows Subsystem for Linux).

## Steps

### 1. Check Your WSL IP Address
You need the internal IP address of your WSL instance.

- Open your **WSL Terminal** (e.g., Ubuntu, Kali Linux).
- Run the following command:

  ```bash
  ip addr

- Look for the `inet` address under your network interface (usually named `eth0`).
- Example output:

  ```
  3: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 ...
      inet 172.25.112.1/20 brd 172.25.127.255 scope global eth0
  ```
- Save the `inet` IP address (e.g., `172.25.112.1`) — you will need it in later steps.

### 2. Check WSL IP Address from Windows
You also need to find out how Windows sees your WSL network.

- Open **Command Prompt** (`cmd`) in Windows.
- Run:

  ```cmd
  ipconfig
  ```

- Look for the network adapter named something like `vEthernet (WSL)`.
- Example output:

  ```
  Ethernet adapter vEthernet (WSL):

     Connection-specific DNS Suffix  . :
     IPv4 Address. . . . . . . . . . . : 172.25.112.1
     Subnet Mask . . . . . . . . . . . : 255.255.240.0
     Default Gateway . . . . . . . . . : 
  ```
- Save the **IPv4 Address** value (e.g., `172.25.112.1`).

### 3. Setup MySQL User
Now, you will create a new MySQL user that is allowed to connect from your WSL instance.

- Access MySQL server using root:

  ```bash
  mysql -u root -p
  ```

- After entering MySQL prompt, run the following commands:

  ```sql
  CREATE USER 'mysql_wsl_username'@'Your_WSL_IP_Address' IDENTIFIED BY 'mysql_wsl_password';
  GRANT ALL PRIVILEGES ON *.* TO 'mysql_wsl_username'@'Your_WSL_IP_Address' WITH GRANT OPTION;
  FLUSH PRIVILEGES;
  ```

> 🔥 **Important**: Replace `Your_WSL_IP_Address` with the IP address you found in Step 1 (e.g., `172.25.112.1`).

### 4. Update Golang Database Connection (`dsn`)
In your Golang project, update your database connection string (`dsn`) like below:

Original example:

```go
dsn := "mysql_wsl_username:mysql_wsl_password@tcp(localhost:3306)/dbname"
```

You must **replace `localhost` with your WSL IP Address (from Windows)**.

Updated example:

```go
dsn := "mysql_wsl_username:mysql_wsl_password@tcp(172.25.112.1:3306)/dbname"
```

> 🚀 **Note:** Replace `172.25.112.1` with your actual WSL IP found in Step 2.

Happy coding with Golang and WSL! 🚀
