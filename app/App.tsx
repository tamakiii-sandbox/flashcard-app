import { StatusBar } from 'expo-status-bar';
import React, { useState } from 'react';
import { StyleSheet, Text, View, TouchableOpacity, SafeAreaView } from 'react-native';
import { FlashcardList } from './src/components/FlashcardList';

export default function App() {
  const [showAllCards, setShowAllCards] = useState(false);

  return (
    <SafeAreaView style={styles.safeArea}>
      <View style={styles.container}>
        {!showAllCards ? (
          <>
            <Text style={styles.title}>Flashcard App</Text>
            <Text style={styles.subtitle}>Improve your learning with flashcards!</Text>
            <TouchableOpacity 
              style={styles.button} 
              onPress={() => setShowAllCards(true)}
              testID="show-all-button"
            >
              <Text style={styles.buttonText}>Show All Flashcards</Text>
            </TouchableOpacity>
          </>
        ) : (
          <>
            <TouchableOpacity 
              style={styles.backButton} 
              onPress={() => setShowAllCards(false)}
              testID="back-button"
            >
              <Text style={styles.backButtonText}>← Back</Text>
            </TouchableOpacity>
            <FlashcardList />
          </>
        )}
        <StatusBar style="auto" />
      </View>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  safeArea: {
    flex: 1,
    backgroundColor: '#f5f5f5',
  },
  container: {
    flex: 1,
    padding: 20,
  },
  title: {
    fontSize: 28,
    fontWeight: 'bold',
    textAlign: 'center',
    marginTop: 40,
    marginBottom: 10,
  },
  subtitle: {
    fontSize: 16,
    textAlign: 'center',
    marginBottom: 40,
    color: '#666',
  },
  button: {
    backgroundColor: '#4a86e8',
    padding: 15,
    borderRadius: 8,
    alignItems: 'center',
    marginVertical: 10,
  },
  buttonText: {
    color: 'white',
    fontSize: 16,
    fontWeight: 'bold',
  },
  backButton: {
    padding: 10,
    marginBottom: 10,
  },
  backButtonText: {
    fontSize: 16,
    color: '#4a86e8',
    fontWeight: 'bold',
  },
});
