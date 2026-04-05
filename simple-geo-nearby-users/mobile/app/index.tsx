import { router } from 'expo-router';
import { useState } from 'react';
import { View, Text, Pressable, TextInput, Alert, Button, StyleSheet } from 'react-native'

export default function ModalScreen() {
  const [username, setUsername] = useState('')

  const handleOnSubmit = () => {
    if (!username) {
      Alert.alert('Entre com o nome do usuário')
      return
    }
    if (username.length > 20) {
      Alert.alert('Nome muito longo, entre com um nome menor que 20 caracteres.')
      return
    }

    router.push('/map')
  }

  return (
    <View style={styles.container}>
      <Text style={styles.headerText}>
        Olá, seja bem vindo(a)
      </Text>
      <View style={styles.form}>
        <Text style={styles.formLabel}>
          Entre com o nome de usuário
        </Text>
        <TextInput
          style={styles.formInput}
          value={username}
          onChangeText={v => setUsername(v)} />
        <Pressable
          style={styles.formButton}
          onPress={handleOnSubmit}>
          <Text style={styles.formButtonText}>
            Entrar
          </Text>
        </Pressable>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    height: '100%',
    width: '100%',
    marginTop: '70%',
    alignItems: 'center',
    gap: 20,
  },
  headerText: {
    fontSize: 20,
    fontWeight: 'bold',
  },
  form: {
    width: '65%',
    gap: 5,
  },
  formLabel: {
  },
  formInput: {
    backgroundColor: '#999',
    width: '100%',
    borderRadius: 5,
  },
  formButton: {
    backgroundColor: '#111',
    width: '100%',
    borderRadius: 5,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 10,
  },
  formButtonText: {
    color: 'white',
  }
})